export interface DraftList {
	id: number;
	Markdown_text?: string
	title: string
}
export interface ApiState {
	DraftList: DraftList[]
}

export interface DraftParameters {
	title: string;
	Markdown_text: string
}

