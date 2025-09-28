import AppCard from './AppCard'
import type { App } from '../types'

export default function AppList({ apps }: { apps: App[] }) {
  if (!apps.length) {
    return <div className="text-center text-gray-600 py-8">No apps yet.</div>
  }

  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
      {apps.map((app) => (
        <AppCard key={app.id ?? app.appId} app={app} />
      ))}
    </div>
  )
}
