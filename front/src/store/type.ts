export interface DraftList {
	id: number;
	markdown_text?: string
	title: string
}
export interface ApiState {
	DraftList: DraftList[]
}

export interface DraftParameters {
	title: string;
	markdown_text: string
}

