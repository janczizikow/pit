export interface NewSubmission {
	name: string;
	class: string;
	mode: string;
	tier: number;
	duration: number;
	video: string;
	build?: string;
}
export interface NewSubmissionRaw {
	name: string;
	class: string;
	mode: string;
	tier: string;
	duration: string;
	video: string;
	build?: string;
}

export interface Submission {
	id: number;
	name: string;
	class: string;
	tier: number;
	mode: string;
	build: string;
	video: string;
	duration: number;
	created_at: string;
	updated_at: string;
}

export interface PaginationMetadata {
	current_page?: number;
	page_size?: number;
	first_page?: number;
	last_page?: number;
	total_records?: number;
}

export interface SubmissionsResponse {
	data: Submission[];
	metadata: PaginationMetadata;
}

export interface APIError {
	status: number;
	errors: { error: string; param: string }[];
}
