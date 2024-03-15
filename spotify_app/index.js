// Grab any variables you need
const react = Spicetify.React;
const reactDOM = Spicetify.ReactDOM;
const {
    URI,
    React: { useState, useEffect, useCallback },
    Platform: { History },
} = Spicetify;

const CONFIG = {
    activeTab: "Main",
    tabs: ["Main","Recent Songs"]
};

// Top Bar Content component
const TopBarContent = (props) => {
    return react.createElement("div", {
        style: {
            display: "flex",
            paddingTop: "15px",
            justifyContent: "flex-start", // Align tabs to the left
            gap: "100px", // Adjust the spacing between tabs
        }
    },
        props.links.map((link) =>
            react.createElement("div", {
                key: link,
                style: {
                    position: "relative",
                }
            },
                react.createElement("a", {
                    className: link === props.activeLink ? 'active' : '',
                    style: {
                        padding: "10px 20px", // Adjust the padding of the tabs
                        borderRadius: "20px", // Add border radius to create rounded edges
                        border: "1px solid white", // Add border with white color
                        color: "white", // Set text color
                        textDecoration: "none", // Remove underline
                    }
                }, link),
            )
        )
    );
};

/*
const fetchTrack = async (uri) => {
    const res = await Spicetify.CosmosAsync.get(`https://api.spotify.com/v1/tracks/${uri.split(':')[2]}`);
    return res.name;
};
*/
function clearsong() {
    return Spicetify.Platform.LocalStorageAPI.clearItem(this.songname)
        .then(() => true) // Resolves to true if the item is successfully cleared
        .catch(() => false); // Resolves to false if there's an error while clearing the item
}
/*used like this: 
clearsong()
    .then(success => {
        if (success) {
            console.log("Item successfully cleared.");
        } else {
            console.error("Failed to clear item.");
        }
    });
*/
function handleLike(){
    //Spicetify.Platform.LocalStorageAPI.setItem(this.songname, { liked: true, skipped: false, song: this.song});

    return "Liked Song";
}
function handleDislike(){
    //Spicetify.Platform.LocalStorageAPI.setItem(this.songname, { liked: false, skipped: false,song: this.song });

    
    return "Disliked Song";
}
function handleSkip(){
    //Spicetify.Platform.LocalStorageAPI.setItem(this.songname, { liked: false, skipped: false, song: this.song });

    return "Skip Song";
}
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
                    display: "flex",
                    justifyContent: "space-between", // This will evenly distribute the items
                    flexWrap: "wrap", // Allow items to wrap to the next line if needed
                },
            }),
            react.createElement("footer", {
                style: {
                    margin: "auto",
                    bottom: 0,
                    left: 0,
                    position: "fixed",
                    textAlign: "center",
                    width: "100%", // Ensure the footer spans the full width
                    paddingTop: "20px", // Add padding at the top for spacing
                    paddingBottom: "40px"
                },
            }, 
                react.createElement("div", {
                    style: {
                        display: "flex",
                        justifyContent: "space-around", // This will evenly distribute the buttons
                        
                    },
                },
                react.createElement("button", {
                    //onClick: () => handleDislike(param1, param2), // Call handleDislike with parameters
                    onClick: handleDislike,
                    style: {
                        backgroundColor: "red", // Change the background color of the button
                        color: "white", // Change the text color of the button
                        border: "none", // Remove the border
                        padding: "10px 20px", // Add padding
                        borderRadius: "5px", // Add border radius
                    }
                }, "Dislike"),
                react.createElement("button", {
                    onClick: handleSkip,
                    style: {
                        backgroundColor: "blue", // Change the background color of the button
                        color: "white", // Change the text color of the button
                        border: "none", // Remove the border
                        padding: "10px 20px", // Add padding
                        borderRadius: "5px", // Add border radius
                    }
                }, "Skip"),
                react.createElement("button", {
                    onClick: handleLike,
                    style: {
                        backgroundColor: "green", // Change the background color of the button
                        color: "white", // Change the text color of the button
                        border: "none", // Remove the border
                        padding: "10px 20px", // Add padding
                        borderRadius: "5px", // Add border radius
                    }
                }, "Like"),)),
                react.createElement(TopBarContent, {
                links: CONFIG.tabs,
                activeLink: CONFIG.activeTab,
            })
        );
    }
}
