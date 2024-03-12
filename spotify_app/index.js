// Grab any variables you need
const react = Spicetify.React;
const reactDOM = Spicetify.ReactDOM;
const {
    URI,
    React: { useState, useEffect, useCallback },
    Platform: { History },
} = Spicetify;

const CONFIG = {
    activeTab: "Tinder",
    tabs: ["Tinder", "Other"]
};

// Load More Icon component
const LoadMoreIcon = (props) => {
    return react.createElement("button", { onClick: props.onClick }, "Load More");
};

// Loading Icon component
const LoadingIcon = () => {
    return react.createElement("div", null, "Loading...");
};

// Top Bar Content component
const TopBarContent = (props) => {
    return react.createElement("div", null,
        props.links.map((link) =>
            react.createElement("a", {
                key: link,
                className: link === props.activeLink ? 'active' : ''
            }, link)
        )
    );
};


const cardList = ["Tinder1", "Other1", "Spot+"];
// The main custom app render function. The component returned is what is rendered in Spotify.
function render() {
    return react.createElement(Grid, { title: "SpotifyPlus" });
}

// Our main component
class Grid extends react.Component {
    constructor(props) {
        super(props);
        Object.assign(this, props);
        this.state = {
            foo: "bar",
            data: "etc"
        };
    }

    render() {
        return react.createElement("section", {
                className: "contentSpacing",
            },
            react.createElement("div", {
                className: "marketplace-header",
            }, react.createElement("h1", null, this.props.title),
            ),
            react.createElement("div", {
                id: "marketplace-grid",
                className: "main-gridContainer-gridContainer",
                "data-tab": CONFIG.activeTab,
                style: {
                    "--minimumColumnWidth": "180px",
                },
            }, [...cardList]),
            react.createElement("footer", {
                style: {
                    margin: "auto",
                    textAlign: "center",
                },
            }, !this.state.endOfList && (this.state.rest ? react.createElement(LoadMoreIcon, { onClick: this.loadMore.bind(this) }) : react.createElement(LoadingIcon)),
            ), react.createElement(TopBarContent, {
                links: CONFIG.tabs,
                activeLink: CONFIG.activeTab,
            })
        );
    }
}
