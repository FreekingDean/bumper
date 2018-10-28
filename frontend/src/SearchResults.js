import React, { Component } from 'react';
import { Table } from 'react-bootstrap';

class SearchResults extends Component {
  constructor(props) {
    super(props);
    this.renderResult = this.renderResult.bind(this);
  }

  renderResult(result, i) {
    return (
      <tr key={i} onClick={() => this.props.handleResultClick(result)}>
        <td><img alt="poster" src={"http://image.tmdb.org/t/p/w92/"+result.poster_path}/></td>
        <td>{result.title}</td>
        <td>{result.media}</td>
      </tr>
    )
  }

  render() {
    if (this.props.isLoading) {
      return (
        "LOADING"
      );
    }

    if (this.props.results.length < 1) {
      return("Nothing Here!");
    }

    return (
      <Table>
        <thead>
          <tr>
            <th>Cover</th>
            <th>Title</th>
            <th>Media Type</th>
          </tr>
        </thead>
        <tbody>
          {this.props.results.map(this.renderResult)}
        </tbody>
      </Table>
    );
  }
}

export default SearchResults;
