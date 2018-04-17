import React from 'react';
import { Form, Input } from 'reactstrap';

import fontawesome from '@fortawesome/fontawesome';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import faSearch from '@fortawesome/fontawesome-free-solid/faSearch';

import './css/sidebar.css';

fontawesome.library.add(faSearch);

export default (props) => {
  return (
    <div className="mr-5">
      <aside>
        <Form className="search-form">
          <Input className="form-control pr-5" type="search" placeholder="Search..." />
          <button className="search-button" type="submit"><FontAwesomeIcon icon="search"/></button>
        </Form>
      </aside>
      <aside className="widget">
        <div className="widget-title">Archives</div>
        <ul>
          <li><a href="#">March 2017</a> (40)</li>
          <li><a href="#">April 2017</a> (08)</li>
          <li><a href="#">May 2017</a> (11)</li>
          <li><a href="#">Jun 2017</a> (21)</li>
        </ul>
      </aside>
      <aside className="widget">
        <div className="widget-title">Tags</div>
        <div className="tagcloud">
          <a href="#">logo</a>
          <a href="#">business</a>
          <a href="#">corporate</a>
          <a href="#">e-commerce</a>
          <a href="#">agency</a>
          <a href="#">responsive</a>
        </div>
      </aside>
    </div>
  );
}
