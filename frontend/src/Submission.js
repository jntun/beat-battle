import React from 'react';
import SCPlayer from './Players/SCPlayer';
import Voter from './Voter';

export default class Submission extends React.Component {
    constructor(props) {
        super(props);
        this.state = {data: props.data};
    }

    componentDidMount() {
        console.log(this.props.id, this.state.data);
    }

    getAppropriatePlayer() {
        switch (this.state.data.type) {
            case 0:
                console.log("soundcloud detected...");
                return <SCPlayer id={"1103179147"}/>;
            default:
                console.log("other detected...");
                return <p className={"error"} id={"player-error"}>Unknown type of submission. Cannot play.</p>
        }
    }

    render() {
        let player = this.getAppropriatePlayer();
        return (
            <div id={this.props.id} className={"submission"}>
                <p>Submission#{this.props.id.substring(0, 8)}</p>
                {player}
                <Voter id={this.props.id} author={this.state.data.author}/>
            </div>
        )
    }
}