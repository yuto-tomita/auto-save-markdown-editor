export interface Draft {
	ID: number;
	Markdown_text?: string
	Title: string
}
export interface ApiState {
	DraftList: Draft[]
}

export interface DraftParameters {
	title: string;
	markdown_text: string
}

