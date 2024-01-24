import React, { memo, useEffect, useRef, useState } from 'react'
import { Button, Modal, Space } from 'antd';
import ProTable from '@ant-design/pro-table';

import i18n from "../../locale/locale";
import { request, waitTime } from '../../utils/utils';
import JsonFormatter from 'react-json-formatter';

const RecordByTaskID = memo(function ({ open, onOK, onCancel, taskInfo: { taskID, taskName } }) {
    const [loading, setLoading] = useState(false);
    const [dataSource, setDataSource] = useState([]);
    const [jsonView, setJsonView] = useState('');
    const [pop, setPop] = useState(false);
    const tableRef = useRef();
    useEffect(() => {
        getData();
        // let id = setInterval(getData, 3000);
        return () => {
            // clearInterval(id);
            setDataSource([]);
            setJsonView('')
        };
    }, [taskID])
    const columns = [
        {
            key: 'getTime',
            title: i18n.t('OVERVIEW.DATE'),
            dataIndex: 'getTime',
            ellipsis: true,
            width: 100
        },
        {
            key: 'getIP',
            title: i18n.t('OVERVIEW.SOURCE_IP'),
            dataIndex: 'getIP',
            ellipsis: true,
            width: 90
        },
        {
            key: 'getResult',
            title: i18n.t('OVERVIEW.RESULT'),
            dataIndex: 'getResult',
            render: (_, v) => {
                return <Button type='link' onClick={() => handleClick(v.getResult)}>{i18n.t('OVERVIEW.VIEW')}</Button>
            },
            ellipsis: true,
            width: 100
        },
    ]
    function handleClick(value) {
        const str = JSON.stringify(JSON.parse(value), null, 2);
        setJsonView(str);
        setPop(true)
    }
    async function getData() {
        await waitTime(300);
        let res = await request('/api/fish/record' + '?taskID=' + taskID);
        let data = res.data;
        setDataSource(data);
        return ({ data: [], success: false, total: 0 });
    }
    return (
        <>
            <Modal
                title={taskName}
                width={500}
                open={open}
                footer={null}
                closeIcon={<></>}
                onOk={onOK}
                onCancel={onCancel}>
                <Space>
                    <ProTable
                        scroll={{
                            x: 'max-content',
                            scrollToFirstRowOnChange: true
                        }}
                        rowKey='recordID'
                        actionRef={tableRef}
                        search={false}
                        columns={columns}
                        onLoadingChange={setLoading}
                        loading={loading}
                        request={getData}
                        pagination={true}
                        dataSource={dataSource}
                        onDataSourceChange={setDataSource}
                    />
                </Space>
            </Modal>
            {
                pop && <JSONView text={jsonView}
                    open={pop}
                    footer={null}
                    closeIcon={<></>}
                    onOk={() => setPop(true)}
                    onCancel={() => setPop(false)}
                />
            }
        </>
    )
})

const JSONView = ({ open, onOK, onCancel, text }) => {
    const jsonStyle = {
        propertyStyle: { color: 'red' },
        stringStyle: { color: 'green' },
        numberStyle: { color: 'darkorange' }
    }

    return <Modal
        open={open}
        footer={null}
        closeIcon={<></>}
        onOk={onOK}
        onCancel={onCancel}>
        <JsonFormatter json={text} tabWith={2} jsonStyle={jsonStyle} />
    </Modal>

}

export default RecordByTaskID;