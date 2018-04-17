import React from 'react';

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
      <h2 className="article__title"><a href="">Learn about Docker!</a></h2>
      <ul className="article__subtitle mb-2">
          <li><FontAwesomeIcon className="mr-2" icon="calendar-alt"/>July 03, 2017</li>
          <li><FontAwesomeIcon className="mr-2" icon="comment-alt"/>3 Comments</li>
          <li><FontAwesomeIcon className="mr-2" icon="tag"/><a href="#">Branding</a>, <a href="#">Design</a></li>
      </ul>
      <a href="#">
        <div className="article__contents">
          <div className="article__contents--image-container">
              <img src="https://userscontent2.emaze.com/images/77789d02-d2d0-4bc9-b976-6bf6d6cbcdb3/84295f963c6adbd26426b822c11fefe6.png"
              alt="Learn about Docker!" className="article__contents--image" />
          </div>
          <p className="article__contents--body text-muted">
            News articles with helpful tips and how-to guides for all things
            related to software development. The main focus of these articles
            is to find better approaches to the software development process.
            News articles with helpful tips and how-to guides for all things
            related to software development. The main focus of these articles
            is to find better approaches to the software development process.
          </p>
        </div>
      </a>
      <hr />
    </article>
  );
}
