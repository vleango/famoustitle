import React from 'react';
import { Link } from 'react-router-dom';
import { Breadcrumb, BreadcrumbItem } from 'reactstrap';

export default class Header extends React.Component {

    render() {
        return (
            <div>
                <Breadcrumb tag="nav">
                    <BreadcrumbItem tag={Link} to="/">Home</BreadcrumbItem>
                    <BreadcrumbItem active tag="span">{this.props.resourceTitle}</BreadcrumbItem>
                </Breadcrumb>
            </div>
        );
    }

}
