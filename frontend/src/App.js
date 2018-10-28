import React, { Component } from 'react';
//import logo from './logo.svg';
import './App.css';
import { Grid, Row, Col } from 'react-bootstrap';
import { Navbar, Nav, NavItem } from 'react-bootstrap';
import { withRouter, Switch, Route, Link } from 'react-router-dom'
import Home from './Home'
import ListSubscriptions from './subscriptions/ListSubscriptions'
import Search from './Search'

class App extends Component {
	constructor(props) {
    super(props);

    this.navigateSearch = this.navigateSearch.bind(this)
    this.navigateSubscriptions = this.navigateSubscriptions.bind(this)
  }

  navigateSearch() {
    this.props.history.push("/media/search")
  }

  navigateSubscriptions() {
    this.props.history.push("/subscriptions")
  }

  render() {
    return (
      <Grid fluid={true}>
        <Navbar>
          <Navbar.Header>
            <Navbar.Brand><Link to="/">Bumper</Link></Navbar.Brand>
          </Navbar.Header>
          <Nav>
            <NavItem href='#' onClick={this.navigateSearch}>Search</NavItem>
            <NavItem href='#' onClick={this.navigateSubscriptions}>View Subscriptions</NavItem>
          </Nav>
        </Navbar>
        <Row>
          <Col xs={12}>
						<Switch>
							<Route exact path='/' component={Home}/>
              <Route path="/media/search" component={Search}/>
              <Route path="/subscriptions" component={ListSubscriptions}/>
						</Switch>
					</Col>
        </Row>
      </Grid>
    );
  }
}

export default withRouter(App);
