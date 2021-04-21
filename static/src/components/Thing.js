
import React from 'react';

class Thing extends React.Component {
    render() {
        return <div className="thingContainer">
                <header>Things Form</header>
                <p>
                    <label className="nameContainer" htmlFor="name">Name:&nbsp;</label><input type="text" className="nameValue" id="name" value={this.props.name} />
                </p>
                <div className="descriptionContainer">
                    <p>Description:</p>
                    <textarea id="description" className="descriptionValue" value={this.props.description} />
                </div>
                <input type="hidden" id="id" value={this.props.id} />
            </div>;
    }
}

export default Thing;
