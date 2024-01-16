package moduleLib

var TaskTableHomeFormat = `<tr style="text-align:center">
                                    <td>%d</td>
                                    <td>%s</td>
                                    <td>%s</td>
                                    <td>%d</td>
                                    <td><span class="badge iq-bg-primary">%s</span></td>
									<td><textarea disabled style="border-width:0;resize:none;font-size:10px;width:100%%;color:gray;">%s</textarea></td>
                                    <td>%s</td>
                                 </tr>`

var TaskTableTasksFormat =
	`<tr style="text-align:center">
		<td>%d</td>
		<td>%s</td>
		<td>%s</td>
		<td>%d</td>
		<td><span class="badge iq-bg-primary" id="taskStatus_%d">%s</span></td>
		<td><textarea disabled style="border-width:0;resize:none;font-size:10px;width:100%%;color:gray;">%s</textarea></td>
		<td>%s</td>
		<td>
			<div class="flex align-items-center list-user-action">
				<a class="btn btn-outline-primary ml-1"  href="#" id="taskControl_%d" onclick="controlTask(%d)">%s</a>
				<a class="btn btn-outline-primary ml-1"  href="#" id="taskApi_%d" onclick="getTaskApi(%d)">API</a>
				<a class="btn btn-outline-primary ml-1"  href="#" id="taskView_%d" onclick="viewTask(%d)">详情</a>
				<a class="btn btn-outline-primary ml-1"  href="#" id="taskDelete_%d" onclick="delTasks(%d)">删除</a>
			</div>
		</td>
	</tr>`

var TaskViewTasksFormat = `<tr style="text-align:center">
                                    <td>%d</td>
                                    <td>%s</td>
                                    <td>%s</td>
                                    <td>%s</td>
                                    <td>
                                       <div class="flex align-items-center list-user-action">
                                          <a class="btn btn-outline-primary ml-1"  href="#" id="recordView_%d" onclick="viewRecord(%d)">详情</a>
                                          <a class="btn btn-outline-primary ml-1"  href="#" id="recordDelete_%d" onclick="delRecord(%d)">删除</a>
                                       </div>
                                    </td>
                                 </tr>`

var TaskUserOptionTasksFormat = `<option value="userModule_%d" id="userModule_%d">%s</option>`
var TaskPublicOptionTasksFormat = `<option value="publicModule_%s" id="publicModule_%s">%s</option>`

var UserCustomModulesFormat = `<div class="col-lg-3 col-md-6">
                                <div class="card card-block card-stretch card-height card-bottom-border-info note-detail">
                                    <div class="card-header d-flex justify-content-between pb-1">
										<a href="#"  id="moduleCustom_%d" onclick="viewModuleCustom(%d)" ><h4 class="card-title" id="moduleCustomTitle_%d">%s</h4></a>
                                    </div>
                                    <div class="card-body rounded">
                                        <p class="mb-3 card-description short" id="moduleCustomData_%d">%s</p>                                                            
                                    </div>
                                    <div class="card-footer text-right">
										<a href="#" class="" id="moduleCustomTime_%d"><i class="las la-calendar mr-2 font-size-20"></i>%s</a>	
                                    </div>
                                </div>
                            </div> `

var AppPublicModulesFormat = `<div class="col-lg-3 col-md-6">
                                <div class="card card-block card-stretch card-height card-bottom-border-info note-detail">
									<div class="card-header d-flex justify-content-between pb-1">
										<a href="#" id="module_%s" onclick="viewModule(%s)"><h4 class="card-title" id="moduleTitle_%s" >%s</h4></a>
                                    </div>
                                    <div class="card-body rounded">
                                        <p class="mb-3 card-description short" id="moduleData_%s">%s</p>                                                            
                                    </div>
                                    <div class="card-footer text-right">
										<a href="#" class="" id="moduleTime_%s"><i class="las la-calendar mr-2 font-size-20"></i>%s</a>	
                                    </div>
                                </div>
                            </div> `

var UserListTableFormat =
	`<tr style="text-align:center">
		<td>%s</td>
		<td>%s</td>
		<td>%s</td>
		<td><span class="badge iq-bg-primary">%s</span></td>
		<td><span class="badge iq-bg-primary">%s</span></td>
		<td>%s</td>
		<td>%s</td>
		<td>%s</td>
		<td>
			<div class="flex align-items-center list-user-action">
				<a class="btn btn-outline-primary ml-1"  href="#" id="userEdit_%s" onclick="editUser(%s)">编辑</a>
				<a class="btn btn-outline-primary ml-1"  href="#" id="userTasks_%s" onclick="tasks(%s)">任务</a>
				<a class="btn btn-outline-primary ml-1"  href="#" id="userModules_%s" onclick="modules(%s)">模块</a>
				<a class="btn btn-outline-primary ml-1"  href="#" id="userDelete_%s" onclick="delUser(%s)">删除</a>
			</div>
		</td>
	</tr>`


var TaskListTableFormat = `<tr style="text-align:center">
                                    <td>%d</td>
                                    <td>%s</td>
                                    <td>%s</td>
                                    <td><span class="badge iq-bg-primary">%s</span></td>
									<td>%s</td>
                                    <td>
                                       <div class="flex align-items-center list-user-action">
                                          <a class="btn btn-outline-primary ml-1"  href="#" id="taskEdit_%d" onclick="editTask(%d)">编辑</a>
                                          <a class="btn btn-outline-primary ml-1"  href="#" id="taskDelete_%d" onclick="delTask(%d)">删除</a>
                                       </div>
                                    </td>
                                 </tr>`

var ModuleListTableFormat = `<tr style="text-align:center">
                                    <td>%d</td>
                                    <td>%s</td>
									<td>%s</td>
                                    <td>
                                       <div class="flex align-items-center list-user-action">
                                          <a class="btn btn-outline-primary ml-1"  href="#" id="moduleEdit_%d" onclick="editModule(%d)">编辑</a>
                                          <a class="btn btn-outline-primary ml-1"  href="#" id="moduleDelete_%d" onclick="delModule(%d)">删除</a>
                                       </div>
                                    </td>
                                 </tr>`