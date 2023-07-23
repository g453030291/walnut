<style>
.demo-ruleForm {
    width: 500px;
    margin: 0 auto;
    padding: 20px 0;
}
textarea {
    height: 200px;
}
</style>

<template>
    <div id="h-form">
        <el-form
                ref="ruleFormRef"
                :model="ruleForm"
                :rules="rules"
                label-width="120px"
                class="demo-ruleForm"
                :size="formSize"
                status-icon
        >
            <el-form-item label="Model">
                <el-select v-model="ruleForm.model" placeholder="please select your model">
                    <el-option label="gpt-3.5-turbo" value="gpt-3.5-turbo" />
<!--                    <el-option label="gpt-4" value="gpt-4" />-->
                    <el-option label="gpt-3.5-turbo-16k" value="gpt-3.5-turbo-16k" />
<!--                    <el-option label="gpt-4-32k" value="gpt-4-32k" />-->
                </el-select>
            </el-form-item>
            <el-form-item label="Id" prop="id">
                <el-input v-model="ruleForm.id"/>
            </el-form-item>
            <el-form-item label="Message" prop="desc">
                <el-input v-model="ruleForm.desc" type="textarea"/>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="submitForm(ruleFormRef)">
                    Create
                </el-button>
                <el-button @click="resetForm(ruleFormRef)">Reset</el-button>
            </el-form-item>
        </el-form>
    </div>
    <div id="result">
        <el-empty description="empty result" />
    </div>
</template>

<script lang="ts" setup>
import {reactive, ref} from 'vue'
import type {FormInstance, FormRules} from 'element-plus'

interface RuleForm {
    model:string
    id:string
    desc: string
}

const formSize = ref('default')
const ruleFormRef = ref<FormInstance>()
const ruleForm = reactive<RuleForm>({
    model:'gpt-3.5-turbo-16k',
    id:'',
    desc: '',
})

const rules = reactive<FormRules<RuleForm>>({
    model: [
        {required: true, message: 'Please input activity form', trigger: 'blur'},
    ],
    id: [
        {required: true, message: 'Please input activity form', trigger: 'blur'},
    ],
    desc: [
        {required: true, message: 'Please input activity form', trigger: 'blur'},
    ],
})

const submitForm = async (formEl: FormInstance | undefined) => {
    if (!formEl) return
    await formEl.validate((valid, fields) => {
        if (valid) {
            console.log('submit!')
        } else {
            console.log('error submit!', fields)
        }
    })
}

const resetForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
}

const options = Array.from({length: 10000}).map((_, idx) => ({
    value: `${idx + 1}`,
    label: `${idx + 1}`,
}))
</script>
