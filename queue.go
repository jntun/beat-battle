package main

import (
	"beat-battle/models"
	"fmt"
)

var submEntryPoint uint = 10
var submHWM uint = 100
var voteEntryPoint uint = 100
var voteHWM uint = 1000

type queueStat struct {
	vote queueIndex
	subm queueIndex
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
	if sess.queueStat.subm.shouldDrain() {
		hwmLog("entry point for submit queue...")
		sess.battle.SubLock.Lock()
		for ; sess.queueStat.subm.lastInsert < sess.queueStat.subm.length; sess.queueStat.subm.lastInsert++ {
			subMsg := sess.submissionQueue[sess.queueStat.subm.lastInsert]
			submission, err := subMsg.ToSubmission()
			if err != nil {
				/* TODO process and see if we want to continue */
				sess.battle.SubLock.Unlock()
				return fmt.Errorf("%v, %s", subMsg, err)
			}
			sess.battle.Submissions[submission.UUID.String()] = *submission
		}
		sess.battle.SubLock.Unlock()
	}
	if sess.queueStat.subm.atHWM() {
		sess.hwmSubm()
	}
	return nil
}

func (sess *Session) drainVoteQueue() error {
	if sess.queueStat.vote.shouldDrain() {
		hwmLog("entry point for vote queue...")
		for ; sess.queueStat.vote.lastInsert < sess.queueStat.vote.length; sess.queueStat.vote.lastInsert++ {
			voteMsg := sess.voteQueue[sess.queueStat.vote.lastInsert]
			if verifyVote(voteMsg) {
				if err := sess.processVote(voteMsg); err != nil {
					return err
				}
			}
		}

	}
	if sess.queueStat.vote.atHWM() {
		sess.hwmVote()
	}
	return nil
}

func (sess *Session) hwmVote() {
	hwmLog("%s", "Flushing vote")
	sess.voteQueue = voteQueue()
	sess.queueStat.vote.lastInsert = uint(0)
	sess.queueStat.vote.length = uint(0)
}

func (sess *Session) hwmSubm() {
	hwmLog("%s", "Flushing submission")
	sess.submissionQueue = submQueue()
	sess.queueStat.subm.lastInsert = uint(0)
	sess.queueStat.subm.length = uint(0)
}

func hwmLog(fmtStr string, args ...interface{}) {
	genLog("Hwm", fmtStr, args...)
}

func submQueue() []models.SubmissionMessage {
	return make([]models.SubmissionMessage, submHWM)
}

func voteQueue() []models.VoteMessage {
	return make([]models.VoteMessage, voteHWM)
}
