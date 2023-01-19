import type { LocalFileInfo } from './LocalFile'

export type UploadFileItem = File | LocalFileInfo | null

export interface UploadTask {
  id: number
  files: UploadFileItem[]
  name: string
  uploadState: 'wait' | 'uploading' | 'pause' | 'done'
  precentage: number
  completedQuantity?: number
}
