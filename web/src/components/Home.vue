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
            <el-form-item label="Content" prop="content">
                <el-input v-model="ruleForm.content" type="textarea"/>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="submitForm(ruleFormRef)">
                    Send
                </el-button>
                <el-button @click="resetForm(ruleFormRef)">Reset</el-button>
            </el-form-item>
        </el-form>
    </div>
    <div>
        <el-text class="mx-1" type="success" v-if="resp">
            <div v-html="resp"></div>
        </el-text>
        <el-empty description="empty result" v-else/>
    </div>
</template>

<script lang="ts" setup>
import {reactive, ref} from 'vue'
import type {FormInstance, FormRules} from 'element-plus'
import axios from "axios";

interface RuleForm {
    model:string
    id:string
    content: string
}

const formSize = ref('default')
const ruleFormRef = ref<FormInstance>()
const ruleForm = reactive<RuleForm>({
    model:'gpt-3.5-turbo-16k',
    id:'',
    content: '',
})

const rules = reactive<FormRules<RuleForm>>({
    model: [
        {required: true, message: 'Please input activity form', trigger: 'blur'},
    ],
    id: [
        {required: true, message: 'Please input activity form', trigger: 'blur'},
    ],
    content: [
        {required: true, message: 'Please input activity form', trigger: 'blur'},
    ],
})

let resp = ref('')

const submitForm = async (formEl: FormInstance | undefined) => {
    if (!formEl) return
    await formEl.validate((valid, fields) => {
        if (valid) {
            axios.post('/api/openai',JSON.stringify(ruleForm)).then((res) => {
                resp.value = res.data.choices[0].message.content.replace(/\n/g, '<br>').replace(/\/\n/g, '<br>')
            }).catch((err) => {
                console.log(err)
            })
            // console.log('submit!')
        } else {
            console.log('error submit!', fields)
        }
    })
}

const resetForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
}

// const options = Array.from({length: 10000}).map((_, idx) => ({
//     value: `${idx + 1}`,
//     label: `${idx + 1}`,
// }))
</script>
