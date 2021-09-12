package main

import (
	"beat-battle/models"
	"fmt"
	"log"
)

var submEntryPoint uint = 10
var submHWM uint = 100
var voteEntryPoint uint = 100
var voteHWM uint = 1000

type queueStat struct {
	voteCount queueIndex
	submCount queueIndex
}

type queueIndex struct {
	// the index of the last place we inserted from
	lastInsert uint

	// all queues have a length of the hwm, however this represents the actual filled length
	length uint

	// entryPoint is the multiple of when we want to start processing
	// and put item(s) until the length into session memory
	entryPoint uint

	/* The highWatermark is the point where we want to dump/drain the queue and re-fill it */
	highWatermark uint
}

/* This is where we determined the threshold we want to actually "drain" the queue at */
/* Drain in this context means process an item at n=lastInsert, make it visible in memory until reaching the i.length*/
func (i *queueIndex) shouldDrain() bool {
	if i.length == 0 {
		return false
	}
	return i.length%i.entryPoint == 0
}

/* Check to see if our queue is at the High Watermark */
func (i *queueIndex) atHWM() bool {
	return i.lastInsert == i.highWatermark
}

func (i *queueIndex) reset() {
	i.lastInsert = uint(0)
	i.length = uint(0)
}

func (sess *Session) drainSubmitQueue() error {
	//hwmLog("%v", sess.queueStat.submCount)
	if sess.queueStat.submCount.shouldDrain() {
		for ; sess.queueStat.submCount.lastInsert < sess.queueStat.submCount.length; sess.queueStat.submCount.lastInsert++ {
			sess.battle.SubLock.Lock()
			subMsg := sess.submissionQueue[sess.queueStat.submCount.lastInsert]
			submission, err := subMsg.ToSubmission()
			if err != nil {
				/* TODO process and see if we want to continue */
				sess.battle.SubLock.Unlock()
				return fmt.Errorf("%v, %s", subMsg, err)
			}
			//log.Printf("draining: %v | ref: %p\n", *submission, submission)
			sess.battle.Submissions[submission.UUID.String()] = *submission
			sess.battle.SubLock.Unlock()
		}
		//hwmLog("sub.len: %d", len(sess.battle.Submissions))
	}
	if sess.queueStat.submCount.atHWM() {
		sess.hwmSubm()
	}
	return nil
}

func (sess *Session) drainVoteQueue() error {
	if sess.queueStat.voteCount.shouldDrain() {
		hwmLog("vote drain: %v", sess.queueStat.voteCount)
		for ; sess.queueStat.voteCount.lastInsert < sess.queueStat.voteCount.length; sess.queueStat.voteCount.lastInsert++ {
			voteMsg := sess.voteQueue[sess.queueStat.voteCount.lastInsert]
			if verifyVote(voteMsg) {
				log.Printf("draining: %v |\n", voteMsg)
				if err := sess.processVote(voteMsg); err != nil {
					return err
				}
			}
		}

	}
	if sess.queueStat.voteCount.atHWM() {
		sess.hwmVote()
	}
	return nil
}

func (sess *Session) hwmVote() {
	hwmLog("%s", "Flushing vote")
	sess.voteQueue = voteQueue()
	sess.queueStat.voteCount.lastInsert = uint(0)
	sess.queueStat.voteCount.length = uint(0)
}

func (sess *Session) hwmSubm() {
	hwmLog("%s", "Flushing submission")
	sess.submissionQueue = submQueue()
	sess.queueStat.submCount.lastInsert = uint(0)
	sess.queueStat.submCount.length = uint(0)
}

func hwmLog(fmtStr string, args ...interface{}) {
	genLog("Hwm", fmtStr, args)
}

func submQueue() []models.SubmissionMessage {
	return make([]models.SubmissionMessage, submHWM)
}

func voteQueue() []models.VoteMessage {
	return make([]models.VoteMessage, voteHWM)
}
