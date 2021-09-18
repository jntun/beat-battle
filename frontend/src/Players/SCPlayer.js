import React from 'react';

export default class SCPlayer extends React.Component {
    constructor(props) {
        super(props);

        this.state = {fetched: false,}
    }

    src = () => {
        let url = "https://w.soundcloud.com/player/?url=https%3A//api.soundcloud.com/tracks/" + this.props.id;
        url = appendUrl(url, "autoplay", present(this.props.autoplay, false))
        url = appendUrl(url, "color", present(this.props.color, "#0066CC"));
        url = appendUrl(url, "buying", present(this.props.buying, false))
        url = appendUrl(url, "sharing", present(this.props.sharing, false))
        url = appendUrl(url, "download", present(this.props.download, false))
        url = appendUrl(url, "show_artwork", present(this.props.show_artwork, true))
        url = appendUrl(url, "show_playcount", present(this.props.show_playcount, false))
        url = appendUrl(url, "show_user", present(this.props.show_user, false))
        url = appendUrl(url, "start_track", present(this.props.start_track, 0))
        url = appendUrl(url, "single_active", present(this.props.single_active, false))
        return url;
    }

    componentDidMount() {
        this.setState({track: this.src(), fetched: true})
    }

    render() {
        let player = <div className={"empty-player"}/>
        if (this.state.fetched)
            player = <iframe width="50%" height="250" scrolling="no" src={this.src()}/>;

        return (
            <div className={"sub-player"}>
                {player}
            </div>
        )
    }

}

function present(val, def) {
    return val ? val : def;
}

function appendUrl(url, str, val) {
    return url += "&" + str + "=" + val;
}
