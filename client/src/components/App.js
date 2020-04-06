import React from 'react';

import api from '../api';

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = { status: 'N/A' };
  }

  componentDidMount() {
    api.get('/')
      .then((res) => this.setState({ status: res.data }));
  }

  render() {
    return <div>SERVER STATUS: [{this.state.status}]</div>;
  }
}

export default App;
