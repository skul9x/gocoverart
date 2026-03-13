export namespace backend {
	
	export class CoverResult {
	    url: string;
	    preview: string;
	    rank: number;
	
	    static createFrom(source: any = {}) {
	        return new CoverResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.preview = source["preview"];
	        this.rank = source["rank"];
	    }
	}
	export class TrackMetadata {
	    filePath: string;
	    title: string;
	    artist: string;
	    album: string;
	    hasCover: boolean;
	    coverData: string;
	
	    static createFrom(source: any = {}) {
	        return new TrackMetadata(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.filePath = source["filePath"];
	        this.title = source["title"];
	        this.artist = source["artist"];
	        this.album = source["album"];
	        this.hasCover = source["hasCover"];
	        this.coverData = source["coverData"];
	    }
	}

}

