export interface ApiState {
	DraftList: {
		id?: number
		markdown_text?: string
		user_id?: number
	}
}

export interface DraftParameters {
	title: string;
	markdown_text: string
}