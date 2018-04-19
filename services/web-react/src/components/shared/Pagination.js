import React from 'react';
import { forEach, range } from 'lodash';
import { Pagination, PaginationItem, PaginationLink } from 'reactstrap';

export default (props) => {
  const {currentPage, totalPages} = props;
  if(currentPage > totalPages) {
    return;
  }

  const pageBoxes = 5;
  let selectedPagePosition = currentPage % pageBoxes;

  if(selectedPagePosition === 0) {
    selectedPagePosition = pageBoxes;
  }

  let paginationItems = [];
  forEach(range(1, pageBoxes + 1), (index) => {
    let pageNumber = currentPage - selectedPagePosition + index;
    paginationItems.push(
      <PaginationItem key={"pagination-item-" + index} active={ selectedPagePosition === index }>
        <PaginationLink href="#">
          { pageNumber }
        </PaginationLink>
      </PaginationItem>
    );
  });

  return (
    <Pagination>
      <PaginationItem disabled={currentPage === 1}>
        <PaginationLink previous href="#" />
      </PaginationItem>
      { paginationItems }
      <PaginationItem disabled={currentPage === totalPages}>
        <PaginationLink next href="#" />
      </PaginationItem>
    </Pagination>
  );
}
