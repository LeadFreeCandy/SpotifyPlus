package internal

type Scope string

const (
	// Images
	UgcImageUpload Scope = "ugc-image-upload"
	// Spotify Connect
	UserReadPlaybackState    Scope = "user-read-playback-state"
	UserModifyPlaybackState  Scope = "user-modify-playback-state"
	UserReadCurrentlyPlaying Scope = "user-read-currently-playing"
	// Playback
	AppRemoteControl Scope = "app-remote-control"
	Streaming        Scope = "streaming"
	// Playlists
	PlaylistReadPrivate       Scope = "playlist-read-private"
	PlaylistReadCollaborative Scope = "playlist-read-collaborative"
	PlaylistModifyPrivate     Scope = "playlist-modify-private"
	PlaylistModifyPublic      Scope = "playlist-modify-public"
	// Follow
	UserFollowModify Scope = "user-follow-modify"
	UserFollowRead   Scope = "user-follow-read"
	// Listening History
	UserReadPlaybackPosition Scope = "user-read-playback-position"
	UserTopRead              Scope = "user-top-read"
	UserReadRecentlyPlayed   Scope = "user-read-recently-played"
	// Library
	UserLibraryModify Scope = "user-library-modify"
	UserLibraryRead   Scope = "user-library-read"
	// Users
	UserReadEmail   Scope = "user-read-email"
	UserReadPrivate Scope = "user-read-private"
	// Open Access
	UserSoaLink            Scope = "user-soa-link"
	UserSoaUnlink          Scope = "user-soa-unlink"
	UserManageEntitlements Scope = "user-manage-entitlements"
	UserManagePartner      Scope = "user-manage-partner"
	UserCreatePartner      Scope = "user-create-partner"
)
