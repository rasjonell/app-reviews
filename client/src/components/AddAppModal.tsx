import { useState, useEffect, useMemo } from 'react'
import { useAddAppMutation } from '../hooks/useApps'
import { useNavigate } from '@tanstack/react-router'

export default function AddAppModal({
  open,
  onClose,
}: {
  open: boolean
  onClose: () => void
}) {
  const [appId, setAppId] = useState('')
  const navigate = useNavigate()
  const addMutation = useAddAppMutation()

  useEffect(() => {
    if (!open) setAppId('')
  }, [open])

  const value = useMemo(() => appId.trim(), [appId])
  const disabled = useMemo(
    () => !value || addMutation.isPending,
    [value, addMutation],
  )

  const onSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!appId.trim()) return
    try {
      await addMutation.mutateAsync({ appId: value })
      onClose()
      navigate({ to: '/reviews/$appId', params: { appId: value } })
    } catch {
      // noop, already handled by mutation.error
    }
  }

  if (!open) return null

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/40 p-4">
      <div className="w-full max-w-sm rounded-xl bg-white shadow-xl p-5">
        <div className="flex items-center justify-between mb-3">
          <h2 className="text-lg font-semibold text-gray-900">Add a New App</h2>
          <button
            className="text-gray-500 hover:text-gray-700 cursor-pointer"
            onClick={onClose}
            aria-label="Close"
          >
            ✕
          </button>
        </div>
        <form onSubmit={onSubmit} className="space-y-3">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              App ID
            </label>
            <input
              value={appId}
              onChange={(e) => setAppId(e.target.value)}
              className="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
              placeholder="e.g. 447188370"
            />
          </div>
          {addMutation.isError && (
            <div className="text-sm text-rose-700 bg-rose-50 border border-rose-200 rounded-md p-2">
              {addMutation.error?.message}
            </div>
          )}
          <div className="flex items-center justify-end gap-2 pt-2">
            <button
              type="button"
              onClick={onClose}
              className="cursor-pointer px-4 py-2 rounded-md border border-gray-300 text-gray-700 hover:bg-gray-50"
            >
              Cancel
            </button>
            <button
              type="submit"
              disabled={disabled}
              className="cursor-pointer px-4 py-2 rounded-md bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-60"
            >
              {addMutation.isPending ? 'Adding…' : 'Add App'}
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}
