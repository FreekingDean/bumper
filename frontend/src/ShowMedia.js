import React, { Component } from 'react';
import { Panel, Button } from 'react-bootstrap';
import { Row, Col } from 'react-bootstrap';

class ShowMedia extends Component {
  render() {
    return (
      <Panel>
        <Panel.Heading>{this.props.media.title}</Panel.Heading>
        <Panel.Body>
          <Row>
            <Col xs={2}>
              <img alt="poster" src={"http://image.tmdb.org/t/p/w92/"+this.props.media.poster_path}/>
            </Col>
            <Col xs={4}><p>{this.props.media.overview}</p></Col>
          </Row>
        </Panel.Body>
        <Panel.Footer><Button onClick={this.props.handleBack}>Back</Button><Button onClick={() => this.props.handleSubscribe(this.props.media)}>+</Button></Panel.Footer>
      </Panel>
    )
  }
}

export default ShowMedia;
