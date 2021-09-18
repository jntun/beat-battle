export const baseEndpoint = 'http://localhost:8000/battle/';

export class userMessage {
    constructor(id, username) {
        this.id = id;
        this.username = username;
    }

    id
    username
}

export class submissionMessage {
    constructor(id, author, resource, type) {
        this.id = id;
        this.author = author;
        this.resource = resource;
        this.type = type;
    }

    id
    author
    resource
    type
}

export class voteMessage {
    constructor(usr, sub) {
        this.user = usr;
        this.submission = sub;
    }

    user
    submission
}