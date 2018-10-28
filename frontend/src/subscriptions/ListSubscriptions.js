import React, { Component } from 'react';
import { Table } from 'react-bootstrap';

class ListSubscriptions extends Component {
  constructor(props) {
    super(props);
    this.renderSubscriptionRow = this.renderSubscriptionRow.bind(this);

    this.state = {
      subscriptions: []
    }
  }

  componentDidMount() {
    fetch(
      `http://127.0.0.1:3001/api/subscriptions`
    ).then((r) => r.json()).then(
      (result) => {
        this.setState({
          subscriptions: result.subscriptions,
        })
      },
      (error) => {
        console.log("ERROR!");
        console.log(error);
      }
    )
  }

  renderSubscriptionRow(subscription, i) {
    console.log(subscription)
    return(
      <tr key={i}>
        <td>{subscription.title}</td>
        <td>{subscription.downloads.length}</td>
        <td>{subscription.disk_versions}</td>
      </tr>
    )
  }

  render() {
    return (
      <Table striped>
        <thead>
          <tr>
            <th>Title</th>
            <th>Downloads</th>
            <th>Disk Versions</th>
          </tr>
        </thead>
        <tbody>
          {this.state.subscriptions.map(this.renderSubscriptionRow)}
        </tbody>
      </Table>
    )
  }
}

export default ListSubscriptions;
