import React, {Fragment} from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';

import {
  Collapse,
  Navbar,
  NavbarToggler,
  NavbarBrand,
  Nav,
  UncontrolledDropdown,
  DropdownToggle,
  DropdownMenu,
  DropdownItem } from 'reactstrap';

import { startLogout } from '../../../actions/auth';

export class ProfileHeader extends React.Component {

  constructor(props) {
    super(props);

    this.state = {
      isOpen: false
    };
  }
  toggle = (e) => {
    this.setState({
      isOpen: !this.state.isOpen
    });
  }

  onLogout = (e) => {
    this.props.startLogout()
  }

  render() {
    return (
      <Navbar color="faded" light expand="md">
        <NavbarBrand tag={Link} to="/">TechWriter</NavbarBrand>
        <NavbarToggler onClick={this.toggle} />
        <Collapse isOpen={this.state.isOpen} navbar>
          <Nav className="ml-auto" navbar>
            <UncontrolledDropdown nav inNavbar>

              { this.props.isAuthenticated ? (
                <Fragment>
                  <DropdownToggle nav caret>
                    { `Hi, ${this.props.firstName}` }
                  </DropdownToggle>
                  <DropdownMenu right>
                    <DropdownItem>
                      Profile
                    </DropdownItem>
                    <DropdownItem>
                      Articles
                    </DropdownItem>
                    <DropdownItem divider />
                    <DropdownItem onClick={this.onLogout}>
                      Logout
                    </DropdownItem>
                  </DropdownMenu>
                </Fragment>

              ) : (
                <Fragment>
                  <DropdownToggle nav caret>
                    Account
                  </DropdownToggle>
                  <DropdownMenu right>
                    <DropdownItem tag={Link} to={'/login'}>
                      Login
                    </DropdownItem>
                    <DropdownItem>
                      Register
                    </DropdownItem>
                  </DropdownMenu>
                </Fragment>
              ) }

            </UncontrolledDropdown>
          </Nav>
        </Collapse>
      </Navbar>
    );
  }

}

const mapStateToProps = (state) => ({
  isAuthenticated: !!state.auth.token,
  firstName: state.auth.firstName,
  lastName: state.auth.lastName
});

const mapDispatchToProps = (dispatch) => ({
  startLogout: () => dispatch(startLogout())
});

export default connect(mapStateToProps, mapDispatchToProps)(ProfileHeader);
