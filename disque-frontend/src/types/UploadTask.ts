export interface UploadTask {
  file: File[] | null
  name: string
  uploadState: 'wait' | 'uploading' | 'pause' | 'done'
  precentage: number
  completedQuantity?: number
}
