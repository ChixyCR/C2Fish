import React from 'react';
import { ModalForm, ProFormGroup, ProFormText } from '@ant-design/pro-form';
import { request } from "../../utils/utils";
import i18n from "../../locale/locale";

const init = {
    taskName: '',
    taskData: '',
    taskStatus: 1,
    moduleID: 1,
    fishUrl: '',
    fishFile: '',
};

function GenerateFish(props) {
    const initValues = getInitValues();
    async function onFinish(form) {
        let formData = {
            ...form,
            moduleID: 1
        };
        let basePath = location.origin + location.pathname + 'api/fish/';
        request(basePath + 'create', formData).then(res => {
            console.log('create fish task res', res)
            if (res.data.status === 'error') {
                console.log('创建失败')
            }
            props.onVisibleChange(false)
        }).catch();
    }

    function getInitValues() {
        return { ...init };
    }

    return (
        <ModalForm
            modalProps={{
                destroyOnClose: true,
                maskClosable: false,
            }}
            initialValues={initValues}
            onFinish={onFinish}
            submitter={{
                render: (_, elems) => elems.pop()
            }}
            {...props}
        >
            <ProFormGroup>
                <ProFormText
                    width="md"
                    name="taskName"
                    label={i18n.t('GENERATOR.FISH_TASK_NAME')}
                    rules={[{
                        required: true
                    }]}
                />
                <ProFormText
                    width="md"
                    name="taskData"
                    label={i18n.t('GENERATOR.FISH_TASK_DATA')}
                    min={1}
                    max={65535}
                    rules={[{
                        required: true
                    }]}
                />
            </ProFormGroup>
            <ProFormGroup>
                <ProFormText
                    width="md"
                    name="taskStatus"
                    label={i18n.t('GENERATOR.FISH_TASK_STATUS')}
                    rules={[{
                        required: true
                    }]}
                />
                <ProFormText
                    width="md"
                    name="fishUrl"
                    label={i18n.t('GENERATOR.FISH_TASK_URL')}
                    rules={[{
                        required: true
                    }]}
                />
            </ProFormGroup>
        </ModalForm>
    )
}

export default GenerateFish;