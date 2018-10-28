import React, { Component } from 'react';
import { Row, Col, FormControl } from 'react-bootstrap';
import SearchResults from './SearchResults';
import ShowMedia from './ShowMedia'

class Search extends Component {
	constructor(props) {
    super(props);

    this.searchMedia = this.searchMedia.bind(this);
    this.handleSearchInput = this.handleSearchInput.bind(this);
    this.handleSearch = this.handleSearch.bind(this);
    this.handleResultClick = this.handleResultClick.bind(this);
    this.handleBack = this.handleBack.bind(this)
    this.handleSubscribe = this.handleSubscribe.bind(this)

    this.state = {
      searchTerm: '',
      searchResults: [],
      searchTimeout: undefined,
      selectedResult: undefined,
    };
  }

  handleBack(e) {
    this.setState({
      selectedResult: undefined
    })
  }

  handleSearch(e) {
    if (e.key === 'Enter') {
      this.searchMedia();
    }
  }

  handleSearchInput(e) {
    this.setState({
      loadingResults: true
    })
    this.setState({
      searchTerm: e.target.value,
    });
  }

  handleResultClick(result) {
    this.setState({
      selectedResult: result
    })
  }

  handleSubscribe(media) {
    console.log(media)
    fetch(
      `http://127.0.0.1:3001/api/subscriptions`,
      {
        method: "POST",
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({imdb_id: media.id})
      }
    ).then((r) => r.json()).then(
      (result) => {
        console.log(result);
        console.log("Added!")
      },
      (error) => {
        console.log("ERROR!");
        console.log(error);
      }
    )
  }


  searchMedia() {
    fetch(
      `http://127.0.0.1:3001/api/searcher?q=${this.state.searchTerm}`
    ).then((resp) => {
      console.log(resp)
      if (!resp.ok) {
        throw resp
      }
      return resp
    }).then(res => res.json()).then(
        (result) => {
          console.log(result);
          this.setState({
            searchResults: result.media,
            loadingResults: false,
          });
        }
    ).catch(
      (error) => {
        console.log("ERROR!");
        console.log(error);
      }
    );
  }

  render() {
    if (this.state.selectedResult !== undefined) {
      return(
        <ShowMedia media={this.state.selectedResult} handleBack={this.handleBack} handleSubscribe={this.handleSubscribe}/>
      )
    }
    return (
      <Row>
        <Col xs={12}>
          <FormControl
            value={this.state.searchTerm}
            onChange={this.handleSearchInput}
            onKeyPress={this.handleSearch}
            type="text"
            placeholder={`Search`}
          />
          <SearchResults isLoading={this.state.loadingResults} results={this.state.searchResults} handleResultClick={this.handleResultClick}/>
        </Col>
      </Row>
    );
  }
}
export default Search;
