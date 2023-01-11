export interface MenuItem{
  text: string
  onClick?(): void
  subItems?: MenuItem[]
}
