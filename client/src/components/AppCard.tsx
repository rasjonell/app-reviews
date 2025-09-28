import { Link } from '@tanstack/react-router'
import type { App } from '../types'

export default function AppCard({ app }: { app: App }) {
  return (
    <div className="group rounded-xl border border-black/10 bg-white shadow-sm hover:shadow-md transition-shadow p-4 flex flex-col gap-2">
      <div className="flex items-center justify-between">
        <h3 className="text-lg font-semibold text-gray-900 truncate">
          {app.name}
        </h3>
        <span
          className={
            'text-xs px-2 py-1 rounded-full ' +
            (app.enabled
              ? 'bg-emerald-100 text-emerald-700'
              : 'bg-rose-100 text-rose-700')
          }
        >
          {app.enabled ? 'Enabled' : 'Disabled'}
        </span>
      </div>
      <div className="text-sm text-gray-600">
        <div className="flex items-center gap-2">
          <span className="font-mono text-gray-800">App ID:</span>
          <span className="font-mono">{app.appId}</span>
        </div>
        {app.lastPolled && (
          <div className="text-xs text-gray-500">
            Last Polled: {new Date(app.lastPolled).toLocaleString()}
          </div>
        )}
      </div>
      <div className="mt-2">
        <Link
          to="/reviews/$appId"
          params={{ appId: app.appId }}
          className="inline-flex items-center justify-center px-3 py-2 text-sm font-medium rounded-md bg-blue-600 text-white hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-400"
        >
          View Reviews
        </Link>
      </div>
    </div>
  )
}
