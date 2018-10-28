import React, { Component } from 'react';
import { Row } from 'react-bootstrap';
import { Link } from 'react-router-dom'

class Home extends Component {
  render() {
    return (
      <Row>
        <Link className="btn btn-primary" to="/subscriptions/new">Add Subscription</Link>
      </Row>
    )
  }
}

export default Home;
