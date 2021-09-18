import {baseEndpoint} from "./Models";
import {userMessage, voteMessage} from "./Models";
import React from 'react';
import axios from 'axios';

export default class Voter extends React.Component {
    doVote = () => {
        let msg = new voteMessage(
            new userMessage("362cfe38-4f4e-428d-aa2f-7bae392d9a99", "test_user"),
            this.props.id
        )
        let jsonMsg = JSON.stringify(msg);
        console.log("sending:", jsonMsg);
        axios.post(baseEndpoint + 'vote', msg)
            .then((resp) => {
                if (resp.status !== 200) {
                    console.log("failed vote msg:", resp)
                }
            });
    }

    render() {
        return (
            <button className={"vote"} onClick={this.doVote}>Vote</button>
        )
    }
}