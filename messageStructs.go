/* This file contains the message stuctures used in the bitTorrent Protocall.
 * Final Project CPE469.
 */

package main

type info_dictionary struct { /* The info on a file and the hashed peices. */
	name         string /* Name of the file. */
	length       int    /* Length of the file in bytes. */
	md5_checksum string /* 32 character hex string of the MD5 checksum of the file. */
	piece_lenth  int    /* Number of bytes in each peice. */
	pieces       string /* String of the concatenation of all 20-byte SHA1 hashes, one per peice. */
	private      bool   /* If 1, client may get peers ONLY from the tracker. */
}

type meta_info struct { /* Holds the metadata on a torrent. */
	info          info_dictionary /* See structure. */
	announce      string          /* The announce URL of the tracker. */
	announce_list string          /* NOT IMPLEMENTED */
	creation_date string          /* The creation time of the torrent since the epoch. */
	comment       string          /* Optional comment by the author. */
	created_by    string          /* Name and version of the program used to create the torrent. */
	encoding      string          /* The encoding type of the .torrent. */
}

type get_request struct { /* A client makes a get request to the tracker to initiate a download. */
	info_hash  string                /* 20-byte SHA1 hash of the value of the info key from meta info. */
	peer_id    string                /* 20-byte string for the unique ID of the client. */
	port       chan tracker_response /*Channel the client is listening on. */
	uploaded   int                   /* Total amount uploaded in bytes. */
	downloaded int                   /* Total amount downloaded in bytes. */
	left       int                   /* Number of bytes needed until 100% downloaded. */
	compact    bool                  /* If 1, use the shortened peers list. */
	event      string                /* Either started, stopped, or completed. */
	ip         string                /* NOT IMPLEMENTED */
	numwant    int                   /* Number of requested peers. May not be honored. */
	key        string                /* NOT IMPLEMENTED */
	trackerid  string                /* If a previous announce contained a tracker id, it should be set here. */
}

type tracker_response struct { /* The tracker repsonds with a pass or fail, with the list of peers upon pass. */
	failure_reason string /* Human readable message of why the request failed. */
	warning        string /* Warning shown like an error but request passes. */
	interval       int    /* Ammount of time in seconds between announces. */
	min_interval   int    /* NOT IMPLEMENTED */
	tracker_id     string /* String that the client should send back on its next announcements. */
	complete       int    /* Number of peers with the complete file (seeders). */
	incomplete     int    /* Number of peers with an incomplete file (leechers). */
	peers          []peer /* We implement the compact form of peerlists. */
}

type handshake struct { /* A handshake is required to be the first messag sent by the client. */
	pstrlen   string  /* Lenth of pstr. Max 255. 19 for BitTorrent. */
	pstr      string  /* String identifier of the protocol. "BitTorrent protocol" */
	reserved  [8]byte /* Reserved for future implementations. */
	info_hash string  /* SHA1 hash of the info key in the metainfo file. Same as tracker requests. */
	peer_id   string  /* Unique ID for the client. Our implementation follows the Azureus-style encoding. */
}

type other_message struct { /* Other messages such as request, piece. */
	offset     int                /* Length of the following data. */
	message_id int                /* A number identifier to the type of message. */
	payload    []byte             /* Data. Optional. */
	file_name  string             /*if empty then no file to search for*/
	port       chan other_message /*Where to return the data to*/
}

type file_info struct { /* Used by the tracker to store info on the files. */
	name         string /* Name of the file. */
	peer_list    []peer /* A slice of who has the file. */
	num_seeders  int    /* The current number of peers with the complete file. */
	num_leechers int    /* The current number of peers with an incomplete file. */
}

type peer struct { /* Holds the data associated with a peer. */
	id   string             /* Identifier */
	port chan other_message /* Channel that the clien talks to. */
}
