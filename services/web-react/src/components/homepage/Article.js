import React from 'react';
import { Link } from 'react-router-dom';
import { truncate } from 'lodash';
import moment from 'moment';

import fontawesome from '@fortawesome/fontawesome';
import faTag from '@fortawesome/fontawesome-free-solid/faTag';

import './css/articles.css';

fontawesome.library.add(faTag);

export default (props) => {
    return (
        <article className="post">
            <Link to={`/articles/${props.article.id}`}>
                <div className="post-preview col-12  no-gutter">
                    <h2>{props.article.title}</h2>

                    {/*<FontAwesomeIcon className="mr-2 subtitle alignment" icon="tag"/>*/}
                    {/*{props.article.tags.map((tag) => {*/}
                        {/*return [*/}
                            {/*<button key={tag} className="subtitle article__subtitle--tag btn btn-link" onClick={() => props.updateFilter("tag", tag)}>{tag}</button>*/}
                        {/*]*/}
                    {/*})}*/}

                    <p className="mt-4">
                        { truncate(props.article.subtitle, {
                            'length': 350,
                            'omission': '...'
                        }) }
                    </p>

                    <p className="meta">
                        {props.article.author} <i className="link-spacer"></i> { moment(props.article.created_at).format("MMMM Do, YYYY") }
                    </p>
                </div>
            </Link>
        </article>
    );
}
