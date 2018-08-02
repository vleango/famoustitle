import React from 'react';

import './css/public-footer.css';

export class PublicFooter extends React.Component {

    render() {
        return (
            <footer className="bd-footer text-muted">
                <div className="container">
                    <ul className="bd-footer-links">
                        <li><a rel="noopener noreferrer" target="_blank" href="https://github.com/vleango/famoustitle">GitHub</a></li>
                        <li><a rel="noopener noreferrer" target="_blank" href="https://stackoverflow.com/users/1316386/tha-leang">Stack Overflow</a></li>
                        <li><a rel="noopener noreferrer" target="_blank" href="https://www.linkedin.com/in/tha-leang-1b340959">LinkedIn</a></li>
                    </ul>
                    <p>Created by Tha Leang</p>
                </div>
            </footer>
        );
    }
}
