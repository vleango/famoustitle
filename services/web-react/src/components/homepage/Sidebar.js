import React from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import { Form, Input } from 'reactstrap';
import moment from 'moment';
import { map } from 'lodash';

import fontawesome from '@fortawesome/fontawesome';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import faSearch from '@fortawesome/fontawesome-free-solid/faSearch';

import './css/sidebar.css';

fontawesome.library.add(faSearch);

export const Sidebar = (props) => {
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
          {
            props.archives && map(props.archives, (count, date) => {
              return <li key={date}><Link to={`/?date=${date}`}>{moment(date).format('MMMM YYYY')} ({count})</Link></li>
            })
          }
        </ul>
      </aside>
      <aside className="widget">
        <div className="widget-title">Tags</div>
        <div className="tagcloud">
          {
            props.tags && props.tags.map((tag) => {
              return <Link to={`/?tag=${tag}`} key={tag}>{tag}</Link>
            })
          }
        </div>
      </aside>
    </div>
  );
}

const mapStateToProps = (state) => {
	return {
		archives: state.articles.index.archives,
    tags: state.articles.index.tags
	};
};

export default connect(mapStateToProps)(Sidebar);
