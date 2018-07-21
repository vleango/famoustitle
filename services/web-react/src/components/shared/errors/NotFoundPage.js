import React from 'react';
import {Helmet} from "react-helmet";

export default class NotFoundPage extends React.Component {

    render() {
        return (
            <div className='container'>
                <Helmet>
                    <title>Page not found - FamousTitle.com</title>
                </Helmet>

                404 - not found
            </div>
        );
    }

}
