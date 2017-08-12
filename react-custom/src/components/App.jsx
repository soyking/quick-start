import React from 'react'

class App extends React.Component {
    render() {
        return (
            <div>
                <div>App</div>
                {this.props.children}
            </div>
        );
    }
}

module.exports = App