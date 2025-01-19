package object

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func Test_Compare(t *testing.T) {
	testCases := []struct {
		name string
		inputA,
		inputB client.Object
		output    bool
		outputErr error
	}{
		{
			name: "Pod objects are equal",
			inputA: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "pod1",
					Namespace: "default",
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "container1",
							Image: "nginx",
						},
					},
				},
			},
			inputB: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "pod1",
					Namespace: "default",
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "container1",
							Image: "nginx",
						},
					},
				},
			},
			output:    true,
			outputErr: nil,
		},
		{
			name: "Pod objects are not equal",
			inputA: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "pod1",
					Namespace: "default",
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "container1",
							Image: "nginx:v1",
						},
					},
				},
			},
			inputB: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "pod1",
					Namespace: "default",
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "container1",
							Image: "nginx:v2",
						},
					},
				},
			},
			output:    false,
			outputErr: nil,
		},
		{
			name: "Pod objects are equal but old contains creation timestamp",
			inputA: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:              "pod1",
					Namespace:         "default",
					CreationTimestamp: metav1.Now(),
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "container1",
							Image: "nginx",
						},
					},
				},
			},
			inputB: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "pod1",
					Namespace: "default",
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "container1",
							Image: "nginx",
						},
					},
				},
			},
			output:    true,
			outputErr: nil,
		},
		{
			name: "Pod objects are equal but old container has additional field",
			inputA: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "pod1",
					Namespace: "default",
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "container1",
							Image: "nginx",
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
			inputB: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "pod1",
					Namespace: "default",
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "container1",
							Image: "nginx",
						},
					},
				},
			},
			output:    true,
			outputErr: nil,
		},
		{
			name: "Pod objects are equal but old missing field",
			inputA: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "pod1",
					Namespace: "default",
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "container1",
							Image: "nginx",
						},
					},
				},
			},
			inputB: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "pod1",
					Namespace: "default",
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "container1",
							Image: "nginx",
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
			output:    false,
			outputErr: nil,
		},
	}
	for _, test := range testCases {
		ok, err := IsEqual(test.inputA, test.inputB)
		if ok != test.output {
			t.Errorf("Test %s failed: expected %t, got %t", test.name, test.output, ok)
		}
		if err != test.outputErr {
			t.Errorf("Test %s failed: expected %v, got %v", test.name, test.outputErr, err)
		}
	}
}
