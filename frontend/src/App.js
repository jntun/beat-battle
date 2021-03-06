import {baseEndpoint} from "./Models";
import React from 'react';
import axios from 'axios';
import Submission from "./Submission";
import './App.css';

class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {submissions: [], data: [], gotData: false, maxView: 10}
    }

    getSubmissions = () => {
        axios.get(baseEndpoint + 'submissions')
            .then((response) => {
                if (response.status === 200) {
                    let submissions = Object.keys(response.data);
                    this.setState({submissions: submissions, data: response.data, gotData: true});
                }
            })
            .catch((err) => {
                console.log("Error getting submissions:", err)
            });
        return false
    }

    componentDidMount() {
        this.setState({submissions: this.getSubmissions()})
    }

    render() {
        let submissions = [];
        if (!this.state.gotData) {
            submissions = <p id={"no-submissions"}>No submissions.</p>
        } else {
            for (let i = 0; i < this.state.submissions.length; i++) {
                if (i === this.state.maxView) {
                    break;
                }
                let x = this.state.submissions[i];
                let data = this.state.data[x];
                submissions.push(<Submission key={x} id={x} data={data}/>)
            }
        }


        return (
            <div className="App">
                <header className="App-header">
                    Beat Battle
                </header>
                {submissions}
            </div>
        );
    }
}

export default App;
