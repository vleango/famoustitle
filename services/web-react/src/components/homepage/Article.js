import React from 'react';
import { Link } from 'react-router-dom';
import { truncate } from 'lodash';
import moment from 'moment';

import fontawesome from '@fortawesome/fontawesome';
import faCalendarAlt from '@fortawesome/fontawesome-free-solid/faCalendarAlt';
import faTag from '@fortawesome/fontawesome-free-solid/faTag';
import faCommentAlt from '@fortawesome/fontawesome-free-solid/faCommentAlt';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';

import './css/articles.css';

fontawesome.library.add(faCalendarAlt, faTag, faCommentAlt);

export default (props) => {
    return (
        <article>
            <h2 className="article__title"><Link to={`/articles/${props.article.id}`}>{props.article.title}</Link></h2>
            <ul className="article__subtitle mb-2">
                <li>
                    <FontAwesomeIcon className="mr-2" icon="calendar-alt"/>
                    <button className="btn btn-link article__subtitle--date" onClick={() => props.updateFilter("date", moment(props.article.created_at).format(`YYYY-MM-01`))}>
                        { moment(props.article.created_at).format("MMMM Do, YYYY") }
                    </button>

                </li>
                {/*<li>*/}
                {/*<FontAwesomeIcon className="mr-2" icon="comment-alt"/>*/}
                {/*<Link to={`/articles/${props.article.id}#comments`}>0 Comments</Link>*/}
                {/*</li>*/}
                { props.article.tags && props.article.tags.length > 0 && (
                    <li className="article__subtitle--tags">
                        <FontAwesomeIcon className="mr-2" icon="tag"/>
                        {props.article.tags.map((tag) => {
                            return [
                                <button key={tag} className="article__subtitle--tag btn btn-link" onClick={() => props.updateFilter("tag", tag)}>{tag}</button>
                            ]
                        })}
                    </li>
                ) }
            </ul>
            <Link to={`/articles/${props.article.id}`}>
                <div className="article__contents">
                    { props.article.img_url && (
                        <div className="article__contents--image-container">
                            <img src={props.article.img_url}
                                 alt={props.article.title} className="article__contents--image" />
                        </div>
                    ) }
                    <p className={props.article.img_url ? 'article__contents--body text-muted' : 'text-muted'}>
                        { truncate(props.article.subtitle, {
                            'length': 350,
                            'omission': '...'
                        }) }
                    </p>
                </div>
            </Link>
            <hr />
        </article>
    );
}
