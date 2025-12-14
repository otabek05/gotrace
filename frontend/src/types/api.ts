export interface ApiResponse<T = any> {
  status: string;  // "success" | "error"
  message: string;
  data: T;
}

