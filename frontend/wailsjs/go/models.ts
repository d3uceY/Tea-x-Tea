export namespace main {
	
	export class FileData {
	    fileName: string;
	    fileContent: string;
	
	    static createFrom(source: any = {}) {
	        return new FileData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fileName = source["fileName"];
	        this.fileContent = source["fileContent"];
	    }
	}
	export class FileInfo {
	    name: string;
	    filePath: string;
	    fileSize: string;
	
	    static createFrom(source: any = {}) {
	        return new FileInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.filePath = source["filePath"];
	        this.fileSize = source["fileSize"];
	    }
	}

}

