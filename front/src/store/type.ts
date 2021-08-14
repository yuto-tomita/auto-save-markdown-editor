export interface DraftList {
	ID: number;
	Markdown_text?: string
	Title: string
}
export interface ApiState {
	DraftList: DraftList[]
}

export interface DraftParameters {
	title: string;
	markdown_text: string
}

