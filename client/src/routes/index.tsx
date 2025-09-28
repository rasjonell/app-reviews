import { createFileRoute } from '@tanstack/react-router'
import { useState } from 'react'
import { useAppsQuery } from '@/hooks/useApps'
import AppList from '@/components/AppList'
import AddAppModal from '@/components/AddAppModal'
import HomeSkeleton from '@/components/HomeSkeleton'

export const Route = createFileRoute('/')({
  component: Home,
})

function Home() {
  const { data, isLoading, isError, error, refetch, isRefetching } =
    useAppsQuery()
  const [modalOpen, setModalOpen] = useState(false)

  return (
    <div className="bg-slate-50 h-screen">
      <div className="max-w-6xl mx-auto px-4 py-8">
        <div className="flex items-center justify-between mb-6">
          <div>
            <h1 className="text-2xl font-bold text-gray-900">Apps</h1>
            <p className="text-gray-600">Browse apps and view reviews</p>
          </div>
          <button
            onClick={() => setModalOpen(true)}
            className="cursor-pointer inline-flex items-center gap-2 px-4 py-2 rounded-md bg-blue-600 text-white hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-400"
          >
            + Add a New App
          </button>
        </div>

        {isLoading && <HomeSkeleton />}

        {isError && (
          <div className="flex items-center justify-between gap-3 bg-rose-50 border border-rose-200 text-rose-800 p-3 rounded-md">
            <div>{error?.message}</div>
            <button
              onClick={() => refetch()}
              className="cursor-pointer px-3 py-1 rounded border border-rose-300 text-rose-700 hover:bg-rose-100"
            >
              Retry
            </button>
          </div>
        )}

        {!!data && !isLoading && !isError && (
          <>
            {isRefetching && (
              <div className="text-sm text-gray-500 mb-2">Refreshing…</div>
            )}
            <AppList apps={data} />
          </>
        )}
      </div>

      <AddAppModal open={modalOpen} onClose={() => setModalOpen(false)} />
    </div>
  )
}
