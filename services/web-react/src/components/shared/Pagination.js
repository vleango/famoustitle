import React, { Fragment } from 'react';
import { Link } from 'react-router-dom';
import { forEach, range } from 'lodash';
import { Pagination, PaginationItem, PaginationLink } from 'reactstrap';

import './css/pagination.css'

export default (props) => {
  const {currentPage, totalPages} = props;
  if(currentPage === 0 || currentPage > totalPages) {
    return (
      <Fragment>{''}</Fragment>
    );
  }

  const pageBoxes = 5;
  let selectedPagePosition = currentPage % pageBoxes;

  if(selectedPagePosition === 0) {
    selectedPagePosition = pageBoxes;
  }

  let paginationItems = [];
  forEach(range(1, pageBoxes + 1), (index) => {
    let pageNumber = currentPage - selectedPagePosition + index;
    if (pageNumber <= totalPages) {
      paginationItems.push(
        <PaginationItem key={"pagination-item-" + index} active={ selectedPagePosition === index }>
          <PaginationLink tag={Link} to={`/?currentPage=${pageNumber}`}>
            { pageNumber }
          </PaginationLink>
        </PaginationItem>
      );
    }
  });

  const nextPageSet = currentPage + (pageBoxes - selectedPagePosition) + 1;

  return (
    <Pagination>
      <PaginationItem hidden={currentPage <= pageBoxes }>
        <PaginationLink previous tag={Link} to={`/?currentPage=${currentPage - selectedPagePosition}`} />
      </PaginationItem>
      { paginationItems }
      <PaginationItem hidden={nextPageSet >= totalPages}>
        <PaginationLink next tag={Link} to={`/?currentPage=${nextPageSet}`} />
      </PaginationItem>
    </Pagination>
  );
}
