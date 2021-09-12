import React from 'react';
import axios from 'axios';

export default class Submission extends React.Component {
    constructor(props) {
        super(props);

        this.state = {data: props.data};
    }

    componentDidMount() {
        axios.get("localhost:8000/battle/submission/" + this.props.id)
            .then((response) => {
                if (response.status === 200) {
                    this.setState({data: response.data})
                }
            })
    }

    render() {
        return (
            <div id={this.props.id} className={"submission"}>
                <p>Submission#{this.props.id}</p>
            </div>
        )
    }
}