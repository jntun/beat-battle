import logo from './logo.svg';
import React from 'react';
import axios from 'axios';
import './App.css';

class App extends React.Component {
    constructor(props) {
        super(props);

        this.state = {submissions: false}
    }

    getSubmissions = () => {
        axios.get('http://localhost:8000/battle/submissions')
            .then((response) => {
                if (response.status === 200) {
                    this.setState({submissions: JSON.parse(response.data)})
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
        console.log(this.state.submissions);
        let submissions;
        if (!submissions) {
            submissions = <p id={"no-submissions"}>No submissions.</p>
        } else {
            submissions = this.state.submissions;
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
