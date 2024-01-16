import { Modal } from 'antd'
import React from 'react'

export default function pop({ open, onOK, onCancel, text }) {
    return (
        <Modal
            open={open}
            onCancel={onCancel}
            onOK={onOK}
            footer={null}
        >
            <div>{text}</div>
        </Modal>
    )
}
