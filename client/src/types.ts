export type App = {
  id: number
  appId: string
  name: string
  enabled: boolean
  lastPolled: string
}

export type Review = {
  id: number
  appId: string
  author: string
  title: string
  content: string
  rating: number
  timestamp: string
}

export type AppReviews = {
  app: App
  reviews: Review[]
}
