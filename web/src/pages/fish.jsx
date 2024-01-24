import React, { useCallback, useEffect, useMemo, memo, useState, useRef } from 'react';
import ProTable from '@ant-design/pro-table';
import { Button, Space, Modal, Switch } from 'antd';
import { request, waitTime } from "../utils/utils";
import i18n from "../locale/locale";
import RecordByTaskID from '../components/record/record';
// DO NOT EDIT OR DELETE THIS COPYRIGHT MESSAGE.
console.log("%c By XZB %c https://github.com/XZB-1248/Spark", 'font-family:"Helvetica Neue",Helvetica,Arial,sans-serif;font-size:64px;color:#00bbee;-webkit-text-fill-color:#00bbee;-webkit-text-stroke:1px#00bbee;', 'font-size:12px;');

let ComponentMap = {
    GenerateFish: null,
    Record: null,
    Pop: null
};

export const loadComponent = (component, callback) => {
    let element = null;
    component = component.toLowerCase();
    Object.keys(ComponentMap).forEach(k => {
        if (k.toLowerCase() === component.toLowerCase()) {
            element = k;
        }
    });
    if (!element) return;
    if (ComponentMap[element] === null) {
        import('../components/' + component + '/' + component).then((m) => {
            ComponentMap[element] = m.default;
            callback();
        });
    } else {
        callback();
    }
};

const Operations = ({ recordClick, onUseCode, onChangeStatus, onDelete, taskID, taskStatus, taskName }) => {
    return (<Space>
        <Button type='primary' onClick={() => recordClick('record', true, { taskID, taskName })}>{i18n.t('OVERVIEW.OPERATION_RECORD')}</Button>
        <Button type='primary' onClick={() => onUseCode(taskID)}>{i18n.t('OVERVIEW.OPERATION_CODE')}</Button>
        <Button type='primary' onClick={() => onChangeStatus(taskID)}>{taskStatus != 0 ? i18n.t('EXPLORER.STOP') : i18n.t('EXPLORER.START')}</Button>
        <Button type='primary' danger onClick={() => onDelete(taskID)}>{i18n.t('EXPLORER.DELETE')}</Button>
    </Space>)
}

const TableCom = (({ dataSource, getData }) => {
    const [loading, setLoading] = useState(false);
    const [columnsState, setColumnsState] = useState(getInitColumnsState());
    const [record, setRecord] = useState(false);
    const [pop, setPop] = useState(false);
    const [popText, setPopText] = useState('');
    const [generateFish, setGenerateFish] = useState(false);
    const [taskInfo, setTaskInfo] = useState({
        taskID: null,
        taskName: ''
    });
    const tableRef = useRef();

    let hooksMap = {
        generateFish: setGenerateFish,
        record: setRecord,
        pop: setPop
    };
    const options = {
        show: true,
        density: true,
        setting: true,
    };
    const fishColumns = [
        {
            key: 'taskID',
            title: i18n.t('GENERATOR.FISH_TASK_ID'),
            dataIndex: 'taskID',
            ellipsis: true,
            render: (_, v) => {
                return String(v.taskID).padStart(3, '0');
            },
            width: 60
        },
        {
            key: 'taskName',
            title: i18n.t('GENERATOR.FISH_TASK_NAME'),
            dataIndex: 'taskName',
            ellipsis: true,
            width: 100
        },
        {
            key: 'taskStatus',
            title: i18n.t('GENERATOR.FISH_TASK_STATUS'),
            dataIndex: 'taskStatus',
            ellipsis: true,
            width: 100,
            render: (_, v) => {
                return <Switch checked={v.taskStatus == 0 ? false : true} />
            }
        },
        {
            key: 'taskData',
            title: i18n.t('GENERATOR.FISH_TASK_DATA'),
            dataIndex: 'taskData',
            ellipsis: true,
            width: 100
        },
        {
            key: 'operation',
            title: '',
            dataIndex: 'operation',
            width: 300,
            render: (_, v) => {
                return <Operations {...v}
                    key={v.taskID}
                    onUseCode={onUseCode}
                    onChangeStatus={onChangeStatus}
                    onDelete={onDelete}
                    recordClick={onClickRecord} />
            }
        }
    ];
    function onClickRecord(act, status, taskInfo) {
        setTaskInfo(taskInfo);
        setRecord(status);
    }
    function onChangeStatus(taskID) {
        console.log(taskID, taskStatus)
        request('/api/fish/stop' + '?taskID=' + taskID);
        tableRef.current.reload();
    }
    function onUseCode(taskID) {
        const str = `<script src=${location.protocol + '//'+ location.hostname+':'+location.port }/api/${btoa(String(taskID).padStart(3, '0'))}></script>`;
        handlePopText(str);
    }
    function onDelete(taskID) {

    }
    function handlePopText(value) {
        setPopText(value);
        onMenuClick('pop', true);
    }
    function onMenuClick(act, status) {
        if (hooksMap[act]) {
            setLoading(true);
            loadComponent(act, () => {
                hooksMap[act](status);
                setLoading(false);
            });
            return;
        }
    }
    function toolBar() {
        return (
            <>
                <Button type='primary' onClick={() => onMenuClick('generateFish', true)}>{i18n.t('OVERVIEW.GENERATEFISH')}</Button>
            </>
        )
    }
    function getInitColumnsState() {
        let data = localStorage.getItem(`fishColumnsState`);
        if (data !== null) {
            let stateMap = {};
            try {
                stateMap = JSON.parse(data);
            } catch (e) {
                stateMap = {};
            }
            return stateMap
        } else {
            localStorage.setItem(`fishColumnsState`, JSON.stringify({}));
            return {};
        }
    }
    function saveColumnsState(stateMap) {
        setColumnsState(stateMap);
        localStorage.setItem(`fishColumnsState`, JSON.stringify(stateMap));
    }

    return <>
        <ProTable
            scroll={{
                x: 'max-content',
                scrollToFirstRowOnChange: true
            }}
            actionRef={tableRef}
            rowKey='taskID'
            search={true}
            options={options}
            columns={fishColumns}
            columnsState={{
                value: columnsState,
                onChange: saveColumnsState
            }}
            onLoadingChange={setLoading}
            loading={loading}
            request={getData}
            pagination={true}
            toolBarRender={toolBar}
            dataSource={dataSource}
        // debounceTime={3}
        />
        <RecordByTaskID
            open={record}
            onOK={() => setRecord(true)}
            onCancel={() => setRecord(false)}
            taskInfo={taskInfo} />
        {
            ComponentMap.GenerateFish &&
            <ComponentMap.GenerateFish
                visible={generateFish}
                onVisibleChange={setGenerateFish}
            />
        }
        {
            ComponentMap.Pop &&
            <ComponentMap.Pop
                open={pop}
                text={popText}
                onOK={() => setPop(true)}
                onCancel={() => setPop(false)}
            />
        }

    </>
})

const FishPage = function (props) {
    const [dataSource, setDataSource] = useState([]);
    useEffect(() => {
        getData();
        let id = setInterval(getData, 10000);
        return () => {
            clearInterval(id);
            setDataSource([])
        };
    }, []);

    const getData = useCallback(async function () {
        let res = await request('/api/fish/task');
        let data = res.data;
        setDataSource(data);
        return ({ data: [], success: false, total: 0 });
    }, [])

    return (
        <>
            <TableCom dataSource={dataSource}
                getData={getData}
            />

        </>
    );
}



function Fish(props) {
    let Component = FishPage;
    return (<Component {...props} key={Math.random()} />)
}

export default Fish;