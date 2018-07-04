import React from 'react';
import { Link } from 'react-router-dom';

export class NavBar extends React.Component {

    render() {
        return (
            <ul className="nav navbar sticky-top navbar-light bg-light">
                <li className="nav-item">
                    <Link className="nav-link active nav-title" to={`/`}>FamousTitle</Link>
                </li>
            </ul>
        );
    }
}
