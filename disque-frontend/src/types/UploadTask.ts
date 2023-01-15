export interface UploadTask {
  id: number
  files: File[] | null
  name: string
  uploadState: 'wait' | 'uploading' | 'pause' | 'done'
  precentage: number
  completedQuantity?: number
}
