import React from 'react';

export default class Submission extends React.Component {
    render() {
        return (
            <div id={this.props.id} className={"submission"}>
                <p>submission</p>
            </div>
        )
    }
}