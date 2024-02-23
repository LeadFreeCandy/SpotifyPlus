// Grab any variables you need
const react = Spicetify.React;
const reactDOM = Spicetify.ReactDOM;

// The main custom app render function. The component returned is what is rendered in Spotify.
function render() {
    return react.createElement("div", { title: "My Custom App" });
}

